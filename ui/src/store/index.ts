import { defineStore } from 'pinia'
import { UserService, AuthenticationError } from '@/services/user.service'
import { StorageTokenService, StorageUserService} from '@/services/storage.service'
import router from '@/router'


export const useAuthStore = defineStore('auth',  {
    
  state: () => {
    return {
      _authenticating: false,
      _accessToken: StorageTokenService.getToken(),
      _authenticationErrorCode: 0,
      _authenticationError: '',
      _username: StorageUserService.getUsername()
    }
  },

  getters:{
    loggedIn: (state) => {
      return state._accessToken ? true : false
    },

    authenticationErrorCode: (state) => {
        return state._authenticationErrorCode
    },

    authenticationError: (state) => {
        return state._authenticationError
    },

    authenticating: (state) => {
        return state._authenticating
    },

    token: (state) => {
      return state._accessToken
    },

    username: (state) : string => {
      return state._username
    }
  },

  actions:{
    async login(username: string, password: string) {
      try {
          this._username = username
          this._accessToken = await UserService.login(username, password);
          console.log(username)

          // non va
          // Redirect the user to the page he first tried to visit or to the home view
          //this.$router.push('/');

          window.location.href = '/'

          return true
      } catch (e) {
          if (e instanceof AuthenticationError) {
            this._authenticationError = true
              console.log('loginError', {errorCode: e.errorCode, errorMessage: e.message})
              this._authenticationErrorCode = e.errorCode 
          }
          return false
      }
  },

  logout() {
      UserService.logout()
      router.push('/login')
  }
  }
})


const mutations = {
  loginRequest(state) {
      state.authenticating = true;
      state.authenticationError = ''
      state.authenticationErrorCode = 0
  },

  loginSuccess(state, accessToken) {
      state.accessToken = accessToken
      state.authenticating = false;
  },

  loginError(state, {errorCode, errorMessage}) {
      state.authenticating = false
      state.authenticationErrorCode = errorCode
      state.authenticationError = errorMessage
  },

  logoutSuccess(state) {
      state.accessToken = ''
  }
}