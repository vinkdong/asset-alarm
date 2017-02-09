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
        formatMoney(value) {
            return "¥" + value;
        }
    },
    mounted: function () {
        this.$nextTick(function () {
            this.listView();
        });
    },
    methods: {
        listView: function () {
            this.$http.get("/api/list").then(res=>{
                this.credits = res.body.credits;
            })
        }
    }
});

var add = new Vue({
    el: "#app-add",
    data: {
        title: "Hello add asset",
        credit: 0,
        debit: 0,
        account_date: 1,
        repayment_date: 1,
        balance: 0,
        name:"VinK Bank"
    },
    filters: {

    },
    mounted: function () {

    },
    methods: {
        addItem: function () {
            this.$http.post("/api/item/add", {
                version: "v0.1",
                credit: {
                    "name": "Vink Bank",
                    "icon": "../icon/vink.logo",
                    "credit": 10.000000,
                    "debit": 50.000000,
                    "balance": 10.000000,
                    "account_date": 8,
                    "repayment_date": 0
                }
            }).then(res => {
                console.log(res)
            });
        }
    }
});

Vue.filter("money",function (value,symbol) {
    return symbol+value;
});