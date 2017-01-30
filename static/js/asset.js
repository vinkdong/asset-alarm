var vm = new Vue({
    el: "#app",
    data: {
        title:"Hello Asset Alarm",
        credits:[]
    },
    filters: {

    },
    mounted: function () {
        this.listView();
    },
    methods: {
        listView: function () {
            var self = this;
            this.$http.get("/api/list").then(function (res) {
                self.credits = res.body.credits;
            })
        }
    }
});