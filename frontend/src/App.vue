<template>

  <div>
    <section class="tepih-servis-azuro" v-cloak>

      <header class="header">
        <h1>Azuro</h1>
      </header>

      <label>Dužina u cm
        <input name="length" type="number" min="0" class="new-todo" autocomplete="off" v-model="length" @keyup.enter="addCustomer">
      </label>
  
      <label>Širina u cm
        <input name="width" type="number" min="0" class="new-todo" autocomplete="off" v-model="width">
      </label>
      
      <label>Cena
         <select name="price" type="number" class="new-todo" autocomplete="off" v-model="price">
          <option value=100>100</option>
          <option value=120>120</option>
          <option value=140>140</option>
          <option value=160>160</option>
         </select>
      </label>
      
      <label>Ime i prezime
        <input name="fullname" class="new-todo" autocomplete="off" v-model="fullname" @keyup.enter="addCustomer">
      </label>

      <label>(Grad) Ulica i broj
        <input name="address" class="new-todo" autocomplete="off" v-model="address" @keyup.enter="addCustomer">
      </label>

      <label>Broj telefona
        <input name="telephoneNumber" class="new-todo" autocomplete="off" v-model="telephoneNumber" @keyup.enter="addCustomer">
      </label>

      <button @click="addCustomer">Dodaj u mušterije</button>

      <section class="main" v-show="customers.length">
        <ul class="todo-list">
          <li class="todo" v-for="customer in customers" :key="customer.telephoneNumber">
            <div class="view">
              <label>{{customer.date}}, {{customer.address}}, {{customer.cost}}</label>
              <button class="destroy" @click="removeCustomer(customer)"></button>
            </div>
          </li>
        </ul>
      </section>

      <button @click="saveAsCsv">Sačuvaj fajl</button>

    </section>
  </div>

</template>

<script>

import "./assets/css/app.css"
import "./assets/css/main.css"
import { jsPDF } from "jspdf";

//import Wails from '@wailsapp/runtime';

export default {
  
  name: "app",

  data() {
    return {
      length: 0,
      width: 0,
      price: 0,
      telephoneNumber: "",
      fullname: "",
      address: "",
      customers: [],
    }
  },

  watch: {
    customers: {
      handler: function(customers) {
        window.backend.Customers.SaveList(JSON.stringify(customers, null, 2));
      },
      deep: true
    }
  },

  mounted() {
    window.backend.Customers.LoadList().then((list) => {
      this.customers = JSON.parse(list);
    });
  },

  methods: {
    addCustomer: function() {
      var width = this.width;
      var length = this.length;
      var price = this.price;
      var telephoneNumber = this.telephoneNumber;
      var fullname = this.fullname;
      var address = this.address;
      if (!width || !length || !price || !telephoneNumber || !fullname || !address) {
        return;
      }

      var today = new Date();

      this.customers.push({
        fullname: fullname,
        address: address,
        telephoneNumber: telephoneNumber,
        cost: price * (width * length) / 10000,
        date: String(today.getDate()).padStart(2, '0') + '/' + String(today.getMonth() + 1).padStart(2, '0') + '/' + today.getFullYear()
      });
      this.width = 0;
      this.length = 0;
      this.price = 0;
      this.telephoneNumber = "";
      this.fullname = "";
      this.address = ""
    },

    removeCustomer: function(customer) {
      var index = this.customers.indexOf(customer);
      this.customers.splice(index, 1);
    },

    saveAsCsv: function() {
      // const doc = new jsPDF();
      // doc.text("hello world!", 10, 10);
      // doc.save("a4.pdf");
      window.backend.Customers.LoadList().then((list) => {
        this.customers = JSON.parse(list);
      });

      var doc = new jsPDF();
      this.customers.forEach(function(customer, i){
      doc.text(20, 10 + (i * 10), 
        customer.fullname + ", " + customer.address + ", " + customer.cost + "din" + "\t - " + customer.telephoneNumber);
      });
      doc.save('Musterije.pdf');
    }

  }
};

</script>