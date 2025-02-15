<script>
	
	let buttonProps = {
	class:[$$restProps.class]
	}
</script>
	<button on:click
			on:mouseover
			on:mouseenter
			on:mouseleave
		{...buttonProps}>
			<slot/>
	</button>

<style>
    
	.primary{
        
		background-color: transparent;
        border: none;
        border-radius: 5px;
	}
	
	.sm {
		font-size:1em;
		padding:0.1em;
	}
	.lg {
		font-size:2em
	}
    button:hover {
        background-color: #9ac0f5; 
    }
    
</style>