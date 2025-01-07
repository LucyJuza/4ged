<template>
  <v-container>
    <v-row
      v-for="(item, index) in items"
    >
      <v-col
        :key="index"
        md="6"
      >
        <v-card class="mb-4 d-flex flex-row align-center">
          <div>
            <v-card-title>Partie du {{ item.date }}</v-card-title>
            <v-card-subtitle>Gagnant·e·s : {{ item.winners }}</v-card-subtitle>
          </div>
          <v-img
            :src="item.image"
            class="align-end" />
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      items: [
        {
          date: "dd/mm/yyyy",
          winners: "name",
          image: "https://via.placeholder.com/400x150",
        },
        {
          date: "dd/mm/yyyy",
          winners: "name",
          image: "https://via.placeholder.com/400x150",
        },
      ],
    };
  },
};
</script>

<style scoped>
.v-card {
  border-radius: 10px;
  overflow: hidden;
}
.v-img {
  object-fit: cover;
}
</style>
