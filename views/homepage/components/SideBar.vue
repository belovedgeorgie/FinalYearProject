(%define "sidebar" %)
<template>
  <v-container fluid>
    <v-row justify="center" align="center">
      <v-col cols="12">
        <v-img src="./assets/unilag.svg" align="left" contain height="100"></v-img>
      </v-col>
      <v-col cols="12">
        <v-textarea
          rounded
          filled
          clearable
          single-line
          auto-grow
          placeholder="search"
          height="8"
          width="100px"
        ></v-textarea>
      </v-col>

      <v-row justify="center">
        <v-col md="auto">
          <v-btn flat outlined height="50" width="50">
            <v-icon>mdi-phone</v-icon>
          </v-btn>
        </v-col>
        <v-col md="1"></v-col>
        <v-col md="auto">
          <v-btn flat outlined height="50" width="50">
            <v-icon>mdi-bell-sleep</v-icon>
          </v-btn>
        </v-col>
      </v-row>
    </v-row>

    <v-flex style="max-height: 580px; max-width: 300px" class="overflow-auto">
      <v-row>
        <v-col cols="12" v-for="i in 5" :key="i">
          <v-card color="#385F73" dark>
            <v-card-title class="headline">Unlimited music now</v-card-title>

            <v-card-subtitle>Listen to your favorite artists and albums whenever and wherever, online and offline.</v-card-subtitle>

            <v-card-actions>
              <v-btn text>Listen Now</v-btn>
            </v-card-actions>
          </v-card>
        </v-col>
      </v-row>
    </v-flex>
  </v-container>
</template>
(%end%)