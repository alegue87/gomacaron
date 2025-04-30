const TOKEN_KEY = 'access_token'
const REFRESH_TOKEN_KEY = 'refresh_token'
const API_KEY = 'api_key'
const USERNAME = 'username'
/**
 * Manage the how Access Tokens are being stored and retreived from storage.
 *
 * Current implementation stores to localStorage. Local Storage should always be
 * accessed through this instace.
**/
const StorageTokenService = {
    getToken() {
        return localStorage.getItem(TOKEN_KEY)
    },
    getApiKey() {
        return localStorage.getItem(API_KEY)
    },
    
    saveToken(accessToken) {
        localStorage.setItem(TOKEN_KEY, accessToken)
    },
   
    saveApikey(apiKey) {
        localStorage.setItem(API_KEY, apiKey)
    },

    removeToken() {
        localStorage.removeItem(TOKEN_KEY)
    },

    getRefreshToken() {
        return localStorage.getItem(REFRESH_TOKEN_KEY)
    },

    saveRefreshToken(refreshToken) {
        localStorage.setItem(REFRESH_TOKEN_KEY, refreshToken)
    },

    removeRefreshToken() {
        localStorage.removeItem(REFRESH_TOKEN_KEY)
    }

}

const StorageUserService = {
    getUsername() {
        return localStorage.getItem(USERNAME)
    },

    saveUsername(username) {
        return localStorage.setItem(USERNAME, username)
    },
}

export { StorageTokenService, StorageUserService }