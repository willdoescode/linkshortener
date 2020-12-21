<template>
  <div class="direct">
    <h1>{{ text }}</h1>
  </div>
</template>

<script>
export default {
  name: "Direct",
  async mounted() {
    if ((await this.GetUrl()) !== undefined) {
      window.location.href = await this.GetUrl();
    } else {
      this.text = "No link is associated with that id.";
    }
  },
  data() {
    return {
      text: "Redirecting you now"
    };
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
