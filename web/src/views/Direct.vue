<template>
  <div class="direct">
    <h1>Hello</h1>
  </div>
</template>

<script>
// @ is an alias to /src

export default {
  name: "Direct",
  async mounted() {
    window.location.href = await this.GetUrl();
  },
  methods: {
    GetUrl() {
      return fetch(`http://localhost:8080/${this.$route.params.id}`)
        .then(data => data.json())
        .then(r => r.url);
    }
  }
};
</script>
