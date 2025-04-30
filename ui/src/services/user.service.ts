import ApiService from './api.service'
import { StorageTokenService, StorageUserService } from './storage.service'


class AuthenticationError extends Error {
    errorCode: any
    constructor(errorCode, message) {
        super(message)
        this.name = this.constructor.name
        this.message = message
        this.errorCode = errorCode
    }
}

const UserService = {
    /**
     * Login the user and store the access token to StorageTokenService. 
     * 
     * @returns access_token
     * @throws AuthenticationError 
    **/
    login: async function(username, password) {

        const env = import.meta.env
        const requestData = {
            method: 'post',
            url: env.VITE_APP_MOQUI_API_ENDPOINT+"/login",
            data: {
                grant_type: 'password',
                username: username,
                password: password
            },
            auth: {
                //username: process.env.VUE_APP_CLIENT_ID,
                //password: process.env.VUE_APP_CLIENT_SECRET
            }
        }
        
        try {

            const response = await ApiService.customRequest(requestData)
            console.log(response)
            StorageTokenService.saveToken(response.data.moquiSessioToken)
            StorageUserService.saveUsername(username)
            StorageTokenService.saveRefreshToken(response.data.moquiSessioToken)
            StorageTokenService.saveApikey(response.data.apiKey)
            ApiService.setHeader()
            
            ApiService.mount401Interceptor();

            return response.data.access_token
        } catch (error) {
            throw new AuthenticationError(error.response.status, error.response.data.detail)
        }
    },

    /**
     * Refresh the access token.
     * TODO: finish
    **/
    refreshToken: async function() {
        const refreshToken = StorageTokenService.getRefreshToken()

        const requestData = {
            method: 'post',
            url: "/o/token/", // TODO
            data: {
                grant_type: 'refresh_token',
                refresh_token: refreshToken
            },
            auth: {
                username: process.env.VUE_APP_CLIENT_ID,
                password: process.env.VUE_APP_CLIENT_SECRET
            }
        }

        try {
            const response = await ApiService.customRequest(requestData)

            StorageTokenService.saveToken(response.data.access_token)
            StorageTokenService.saveRefreshToken(response.data.refresh_token)
            // Update the header in ApiService
            ApiService.setHeader()

            return response.data.access_token
        } catch (error) {
            throw new AuthenticationError(error.response.status, error.response.data.detail)
        }

    },

    /**
     * Logout the current user by removing the token from storage. 
     * 
     * Will also remove `Authorization Bearer <token>` header from future requests.
    **/
    logout() {
        // Remove the token and remove Authorization header from Api Service as well 
        StorageTokenService.removeToken()
        StorageTokenService.removeRefreshToken()
        ApiService.removeHeader()
        
        // NOTE: Again, we'll cover the 401 Interceptor a bit later. 
        ApiService.unmount401Interceptor()
    }
}

export default UserService

export { UserService, AuthenticationError }