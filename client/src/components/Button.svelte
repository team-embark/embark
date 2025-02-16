<script>
	export let disabled = false;
	let buttonProps = {
	class:[$$restProps.class]
	}

</script>
	<button on:click
			on:mouseover
			on:mouseenter
			on:mouseleave
            {disabled}
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
            top: 500px;
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
            filter: grayscale(1);
            top: 1200px;
            left: 600px;
        }
        .button5 {
            filter: grayscale(1);
            top: 320px;
            left: 350px;
        }
        .transparentBtn {
            border-width: 0px;
            position:fixed;
            left:20px;
            z-index:30;
            color: black;
            font-size: 1.5rem;
            background-color: transparent;
            
        }
        .transparentBtn1 {
            border-width: 0px;
            position:fixed;
            right:100px;
            top:45px;
            color: black;
            font-size: 1.5rem;
            z-index:20;
            background-color: transparent;
            
        }
    

        .button1:hover, .button2:hover, .button3:hover,.transparentBtn:hover, .transparentBtn1:hover {
    
            transform: scale(1.1);
        }
    
        .image-container:hover {
            background-color: #e0e0e0;
            transform: scale(1.02);
            transition: transform 0.2s ease-in-out, background-color 0.2s ease-in-out;
        }

        button {
            font-family: "Futura", "Trebuchet MS", Arial, sans-serif;
            font-style: italic;
            font-weight: 500;
        }
    </style>
