<template>
  <div class="home">
    <form class="form">
      <label>
        <input
          class="form-input"
          v-model="url"
          type="text"
          placeholder="Enter url."
        />
      </label>
      <button class="sub" v-on:click="postId">Submit</button>
    </form>
    <div class="output">Output: {{ output }}</div>
  </div>
</template>

<style lang="scss">
.home {
  position: absolute;
  top: 50%;
  left: 50%;
  margin-right: -50%;
  transform: translate(-50%, -50%);
  .form {
    display: flex;
    flex-direction: column;
    .form-input {
      font-size: 2em;
      border: 1px solid mediumpurple;
      padding: 0.6em;
      font-weight: lighter;
    }
    .sub {
      font-size: 2em;
      border: 1px solid mediumpurple;
      padding: 0.6em;
      text-align: left;
      color: purple;
      font-weight: bold;
      background: mediumpurple;
    }
  }
}
.output {
  position: absolute;
  font-weight: bold;
  color: darkslateblue;
}
</style>

<script>
import { nanoid } from "nanoid";

export default {
  name: "Home",
  data() {
    return {
      url: "",
      output: window.location.href
    };
  },
  methods: {
    async postId(e) {
      e.preventDefault();
      this.url = this.url.trim();
      if (this.url === "") {
        return;
      }
      if (!this.url.startsWith("https://") && !this.url.startsWith("http://")) {
        this.url = `https://${this.url}`;
      }
      const id = nanoid(4);
      fetch("http://localhost:7000/create", {
        method: "post",
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json"
        },

        body: JSON.stringify({
          id: id,
          url: this.url
        })
      })
        .then(response => response.json())
        .then(r => {
          this.output = `${window.location.href}${r.id}`;
          this.$copyText(this.output);
        });
    }
  }
};
</script>
