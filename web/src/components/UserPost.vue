<template>
    <div id="user-post">
        <div class="card">
            <div class="card-body">
                
            </div>
        </div>
    </div>
</template>
<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
// import store from "../store";
// import axios from "axios";
// import Global from "../global";

@Component
export default class UserPost extends Vue {
    menuUserPost:any=[
        {
            id:1,
            name:"map"
        },
         {
            id:2,
            name:"picture"
        },
         {
            id:3,
            name:"post_detail"
        },
         {
            id:4,
            name:"add_article"
        }
    ];
}
</script>
