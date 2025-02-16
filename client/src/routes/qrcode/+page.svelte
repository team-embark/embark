<script>
  import CenteredBox from "./CenteredBox.svelte";
  import NewTextComponent from "./NewTestComponent.svelte";
  import ButtonComponent from "./[park]/ButtonComponent.svelte";

  let {parkname} = $props();
</script>
<div>
	<p>Thanks for visiting {parkname}!</p>
	<CenteredBox width="60em" height="40em" />
	<NewTextComponent text="Fun fact!" />
	<ButtonComponent label="To Journal" />
	<img src="/YosemiteNationalPark-2.png" alt="YosemiteNationalPark-2" width= "200" class="image"/>
	<style>
		.image {
		  position: fixed;
		  bottom: 450px; /* 20px from bottom */
		  right: 1075px; /* 20px from right */
		}
	  </style>
</div>

<style>
	p {
		color: goldenrod;
		font-family: 'Comic Sans MS', cursive;
		font-size: 2em;
    text-align: center;
	}
</style>
