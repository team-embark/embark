<script>
    export let label = "Click Me";
  </script>
  
  <button class="btn">{label}</button>
  
  <style>
    .btn {
      background-color: forestgreen;
      color: white;
      font-size: 1.5em;
      padding: 20px 40px;
      border: none;
      border-radius: 12px;
      cursor: pointer;
      transition: background-color 0.3s ease;
      position: absolute;
      top: 700px;  /* Moves text up or down */
      left: 1000px; /* Moves text left or right */
    }
  
    .btn:hover {
      background-color: darkgreen;
    }
  </style>
  