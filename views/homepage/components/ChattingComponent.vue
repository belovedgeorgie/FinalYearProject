(%define "chattingComponent"%)
<template>
  <v-container fluid>
    <v-row no-gutters>
      <v-col cols="12">
        <v-row no-gutters>
          <v-col>
            <v-col md="12"></v-col>
            <h2>
              <b>Name Here (Matric@live.unilag.edu.ng)</b>
            </h2>
          </v-col>

          <v-col md="auto" align="right">
            <v-btn fab outline flat depressed>
              <v-icon>mdi-video-outline</v-icon>
            </v-btn>
            <v-btn fab outline flat depressed>
              <v-icon>mdi-account-multiple-plus-outline</v-icon>
            </v-btn>
            <v-btn fab outline flat depressed>
              <v-icon>mdi-phone</v-icon>
            </v-btn>
          </v-col>
        </v-row>
        <v-divider></v-divider>
        <v-col cols="12"></v-col>
      </v-col>

      <v-flex style="max-height: 620px; " class="overflow-auto">
        <v-row>
          <v-col cols="12" align="right">
            <v-chip class="ma-2">Hello</v-chip>
          </v-col>
          <v-col cols="12" align="left">
            <v-chip class="ma-2">typing....</v-chip>
          </v-col>
          <v-col cols="12" align="right">
            <v-chip class="ma-2">Hello</v-chip>
          </v-col>
          <v-col cols="12" align="left">
            <v-chip class="ma-2">typing....</v-chip>
          </v-col>
          <v-col cols="12" align="right">
            <v-chip class="ma-2">Hello</v-chip>
          </v-col>
          <v-col cols="12" align="left">
            <v-chip class="ma-2">typing....</v-chip>
          </v-col>
          <v-col cols="12" align="right">
            <v-chip class="ma-2">Hello</v-chip>
          </v-col>
          <v-col cols="12" align="left">
            <v-chip class="ma-2">typing....</v-chip>
          </v-col>
          <v-col cols="12" align="right">
            <v-chip class="ma-2">Hello</v-chip>
          </v-col>
          <v-col cols="12" align="left">
            <v-chip class="ma-2">typing....</v-chip>
          </v-col>
        </v-row>
      </v-flex>

      <v-col></v-col>
    </v-row>
  </v-container>
</template>
(%end%)