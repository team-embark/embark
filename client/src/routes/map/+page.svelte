


<style>
    :global(body) {
      background-color: #b5e6ff;
    }
  </style>

<script>
  import Button from "./Button.svelte";  
</script>


<div class="image-container">
    <img src="/California2.png" alt="California2" width= "1300"/>  

<Button class="button1" on:click={()=>console.log("Clicked1")}>

    <img src="/YosemiteNationalPark-2.png" alt="YosemiteNationalPark-2" width= "100"/>
    <br>
    Yosemite
</Button>

<Button class="button2" on:click={()=>console.log("Clicked")}>
    <img src="/RedWoodNationalPark-2.png" alt="RedWoodNationalPark-2" width= "100" />
    <br>
    Red Wood
</Button>

<Button class="button3" on:click={()=>console.log("Clicked")}>
    <img src="/JoshuaTree-2.png" alt="JoshuaTree-2" width= "100" />
    <br>
    Joshua Tree
</Button>
   
<Button class="button4" on:click={()=>console.log("Clicked")}>
    <img src="/channelIslands.png" alt="channelIslands" width= "100" />
    <br>
    Channel Islands
</Button>

</div>