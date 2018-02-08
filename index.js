Vue.component("unit-list", {
    template: "#unit-list",
    props: ["state"],
    data: function() { return { searchString: '', userOnly: true } },
    methods: {
        click: function() {
            this.state.reload()
        },
        search: function() {
            this.state.search(this.searchString, this.userOnly);
        },
    },
});

var vm = new Vue({
  el: '#app',
  data: {
    systemState: window.systemState,
  },
});

window.systemState.search("", true)
