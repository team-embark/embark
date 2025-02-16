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
            width: 300px; 
            height: 300px; 
        }
    
        .button1, .button2, .button3, .button4, .button5 {
            position: absolute;
            background-color: transparent;
            border: none;
            border-radius: 5px;
            transition: transform 0.2s ease-in-out, background-color 0.2s ease-in-out;
        }
    
        .button1 {
            top: 200px;
            left: 400px;
        }
    
        .button2 {
            top: 200px;
            left: 100px;
        }
    
        .button3 {
            top: 1000px;
            left: 800px;
        }
    
        .button4 {
            top: 1200px;
            left: 600px;
        }
        .button5 {
            top: 320px;
            left: 350px;
        }

        .button1:hover, .button2:hover, .button3:hover, .button4:hover, .button5:hover {
           
            transform: scale(1.1);
        }
    
      
        .image-container:hover {
            background-color: #e0e0e0;
            transform: scale(1.02);
            transition: transform 0.2s ease-in-out, background-color 0.2s ease-in-out;
        }
    </style>