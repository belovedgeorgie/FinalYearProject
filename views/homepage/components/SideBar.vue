(%define "sidebar" %)
<template>
  <v-container fluid>
    <v-row justify="center" align="center">
      <v-col cols="12">
        <v-img src="./assets/unilag.svg" align="left" contain height="100"></v-img>
      </v-col>

      <v-col md="auto">
        <v-text-field rounded filled clearable placeholder="search for contacts"></v-text-field>
      </v-col>

      <v-row justify="center" align="center">
        <v-col md="auto">
          <v-btn flat outlined height="50" width="50">
            <v-icon>mdi-phone</v-icon>
          </v-btn>
        </v-col>
        <v-col md="auto" offset-md="1">
          <v-btn flat outlined height="50" width="50">
            <v-icon>mdi-bell</v-icon>
          </v-btn>
        </v-col>
      </v-row>
      <v-col cols="12"></v-col>
    </v-row>

    <v-flex fluid style="max-height: 490px; max-width: 300px" class="overflow-y-auto">
      <v-row no-gutters>
        <v-list tile dense three-line>
          <v-list-item-group v-model="item" color="black">
            <v-list-item v-for="i in 10" :key="i">
              <v-list-item-avatar>
                <v-icon large>mdi-account-circle</v-icon>
              </v-list-item-avatar>

              <v-list-item-content>
                <v-list-item-title>matric@live.unilag.edu.ng</v-list-item-title>
                <v-list-item-subtitle>Text Here</v-list-item-subtitle>
              </v-list-item-content>
            </v-list-item>
          </v-list-item-group>
        </v-list>
      </v-row>
    </v-flex>
  </v-container>
</template>
        </v-col>
      </v-row>
    </v-flex>
  </v-container>
</template>
(%end%)

(%define "sidebarData"%)
  data: () => ({
    item: 0,
  })
(%end%)