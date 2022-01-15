<template>
  <div id="app">
    <div id="form">
      <h1>Art Gallery</h1>
      <!-- <form method="POST" @submit.prevent="sendItem()">
        <input type="text" size="50" v-model="paintingitem" placeholder="Enter new item"/>
        <input type="submit" value="Submit"/>
      </form> --->
    </div>
    <div id="list">
      <ul>
        <li v-for="item in paintings" v-bind:key="item">
          {{ item }}
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
import axios from "axios";
const appData = {
  paintings: []
};

export default {
  name: 'App',
  data() {
    return appData;
  },
  mounted: function() {
    getPaintings();
  },
  methods: {
    getPaintings: getPaintings,
    sendItem: sendItem,
  }
};

function getPaintings() {
  axios.get("/api/painting").then((res) => {
    appData.paintings = res.data.list;
  });
}

async function sendItem() {
  const params = new URLSearchParams();
  params.append("item", this.paintingitme);
  await axios
    .post("/api/painting/create", params)
    .then(function () {
      getPaintings();
    })
    .catch(function (error) {
      appData.paintings = [error.message];
    });
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  margin-top: 60px;
}
#form {
  text-align: center;
  margin-top: 60px;
}

</style>
