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
      .image-container {
        position: relative; 
        display: inline-block; 
     }
     .image-container Button {
    
    position: absolute; 
    transform: translate(-50%, -50%); 

    }

	.button1{   
        position: absolute; 
        top: 22%;
        left: 28%;
		background-color: transparent;
        border: none;
        border-radius: 5px;
        aspect-ratio:1

	}
	.button2{   
        position: absolute; 
        top: 20%;
        left: 8%;
		background-color: transparent;
        border: none;
        border-radius: 5px;
        aspect-ratio:1

	}.button3{   
        position: absolute; 
        bottom: -40%;
        right: 35%;
        
		background-color: transparent;
        border: none;
        border-radius: 5px;
        aspect-ratio:1
	}
    button:hover {
        background-color: #9ac0f5; 
    }
    
</style>