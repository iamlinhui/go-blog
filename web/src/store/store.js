import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
    modules: {
        user: {
            name: "",
            token: ""
        }
    }
})
export default store