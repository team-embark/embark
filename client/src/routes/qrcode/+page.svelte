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
	<img src="/JoshuaTree-2.png" alt="Joshua Tree National Park" width= "500" class="image"/>
	<style>
		.image {
		  	display: flex;
    		justify-content: center;        /* Centers content horizontally */
    		align-items: center; 
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
