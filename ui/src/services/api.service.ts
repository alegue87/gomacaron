import axios from 'axios'
import { StorageTokenService } from '@/services/storage.service'


const ApiService = {
    // Stores the 401 interceptor position so that it can be later ejected when needed
    _401interceptor: null,

    init(baseURL) {
        axios.defaults.baseURL = baseURL;
    },

    setHeader() {
        axios.defaults.headers.common["moquiSessionToken"] = StorageTokenService.getToken()
        axios.defaults.headers.common["api_key"] = StorageTokenService.getApiKey()
    },

    removeHeader() {
        axios.defaults.headers.common = {}
    },

    get(resource) {
        return axios.get(resource)
    },

    post(resource, data) {
        return axios.post(resource, data)
    },

    put(resource, data) {
        return axios.put(resource, data)
    },

    delete(resource) {
        return axios.delete(resource)
    },

    /**
     * Perform a custom Axios request.
     *
     * data is an object containing the following properties:
     *  - method
     *  - url
     *  - data ... request payload
     *  - auth (optional)
     *    - username
     *    - password
    **/
    customRequest(data) {
        return axios(data)
    },

    mount401Interceptor() {
        this._401interceptor = axios.interceptors.response.use(
            (response) => {
                return response
            },
            async (error) => {
                if (error.request.status == 401) {
                    if (error.config.url.includes('/login')) {
                        // Refresh token has failed. Logout the user
                        // authStore.dispatch('auth/logout')
                        throw error
                    } else {
                        // Refresh the access token
                        try{
                            // await authStore.dispatch('auth/refreshToken')
                            // Retry the original request
                            return this.customRequest({
                                method: error.config.method,
                                url: error.config.url,
                                data: error.config.data
                            })
                        } catch (e) {
                            // Refresh has failed - reject the original request
                            throw error
                        }
                    }
                }

                // If error was not 401 just reject as is
                throw error
            }
        )
    },

    unmount401Interceptor() {
        // Eject the interceptor
        axios.interceptors.response.eject(this._401interceptor)
    }
}

export default ApiService