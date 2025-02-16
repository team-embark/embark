<script>
    export let width = "40em";
    export let height = "30em";
  </script>
  
  <style>
    :global(body) {
      display: flex;
      justify-content: center;
      align-items: center;
      margin: 0;
      background-color: #f0f0f0;
    }
  
    .box {
      width: var(--width);
      height: var(--height);
      border-radius: 3em;
      background-color: #06402B;
    }
  </style>
  
  <div class="box" style="--width: {width}; --height: {height};"></div>
  