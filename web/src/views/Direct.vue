<template>
  <div class="direct">
    <h1>{{ text }}</h1>
  </div>
</template>

<script>
export default {
  name: "Direct",
  async mounted() {
    let res = await this.GetUrl();
    if (res !== undefined) {
      this.text = `Redirecting you to: ${res}`;
      window.location.href = res;
    } else {
      this.text = `Could not find url associated with the id: "${this.$route.params.id}"`;
    }
  },
  data() {
    return {
      text: "Redirecting you now"
    };
  },
  methods: {
    GetUrl() {
      return fetch(`http://localhost:7000/${this.$route.params.id}`)
        .then(data => data.json())
        .then(r => r.url);
    }
  }
};
</script>
