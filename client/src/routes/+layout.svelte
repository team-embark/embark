<script>
    import Button from "./map/button.svelte";
	let { children } = $props();
</script>


<div class="banner">
  <nav>
	<a href="/">home</a>
	<a href="/">map</a>
    EMBARK
    <a href="/signin">Sign Out</a>
</nav>
</div>

<style>
    .banner {
        background:rgb(193, 122, 67);
        padding: 20px;
        color: white;
        font-size: 1.5rem;
        text-align: center;
        border-radius: 4px;
    }
</style>

{@render children()}

