


<style>
    :global(body) {
      background-color: #b5e6ff;
    }
  </style>

<script>
  import Button from "./Button.svelte";  
</script>


<div class="image-container">

<Button class="button1" on:click={()=>console.log("Clicked")}>

    <img src="/YosemiteNationalPark-2.png" alt="YosemiteNationalPark-2" width= "100" />
</Button>

<Button class="button2" on:click={()=>console.log("Clicked")}>
    <img src="/RedWoodNationalPark-2.png" alt="RedWoodNationalPark-2" width= "100" />
</Button>

<Button class="button3" on:click={()=>console.log("Clicked")}>
    <img src="/JoshuaTree-2.png" alt="JoshuaTree-2" width= "100" />
</Button>
    <img src="/California.png" alt="California Image" width= "1300" />  

</div>