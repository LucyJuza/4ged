<template>
  <v-container>
    <v-row
      v-for="(item, index) in list"
    >
      <v-col
        :key="index"
        md="6"
      >
        <v-card class="mb-4 d-flex flex-row align-center">
          <div class="flex-grow-1">
            <v-card-title>{{ item.title }}</v-card-title>
            <v-card-subtitle>{{ item.subtitle }}</v-card-subtitle>
          </div>
          <v-img
            :src="item.image"
            class="align-end" />
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
defineProps({
  list: Array
})
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
