const store = new Vuex.Store({
    state: {
        isCtrl: false,
        isUserSearch: true,
        searchString: "",
        autoreload: false,
    },
    mutations: {
        setCtrl: function(state, value) {
            state.isCtrl = value
        },
        setUserSearch: function(state, value) {
            state.isUserSearch = value
        },
        setSearchString: function(state, value) {
            state.searchString = value
        },
        toggleUserSearch: function(state) {
            state.isUserSearch = !state.isUserSearch
        },
    }
})

document.onkeyup = function(e) {
    if (e.which == 17) store.commit('setCtrl', false);
}

document.onkeydown = function(e) {
    if (e.which == 17) store.commit('setCtrl', true);

    if (store.state.isCtrl) {
        if (e.which == 85) { // u
            store.commit('toggleUserSearch');
        }
    }
}


Vue.component("unit-list", {
    template: "#unit-list",
    props: ["state"],
    computed: {
        isUserSearch: {
            get: function() { return store.state.isUserSearch },
            set: function(value) { store.commit('setUserSearch', value) }
        },
        searchString: {
            get: function() { return store.state.searchString },
            set: function(value) { store.commit('setSearchString', value) }
        },
    },
    methods: {
        search: function() {
            this.state.search(this.searchString, this.isUserSearch);
        },
    },
});

var vm = new Vue({
    el: '#app',
    data: {
        systemState: window.systemState,
    },
});

var reloadSystemd = function() {
    window.systemState.search(store.state.searchString, store.state.isUserSearch);
}

reloadSystemd()
