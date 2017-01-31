var vm = new Vue({
    el: "#app",
    data: {
        title:"Hello Asset Alarm",
        context: {
            pages:{
                add_item_page: "./add.html"
            }
        },
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

var add = new Vue({
    el: "#app-add",
    data: {
        title:"Hello add asset"
    },
    filters: {

    },
    mounted: function () {

    },
    methods: {
    }
});