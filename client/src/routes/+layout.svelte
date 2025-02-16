<script>
    import Button from "./map/button.svelte";
    import {auth} from "../firebase_yippee";
    import {goto} from "$app/navigation";
	let { children } = $props();

    function logout() {
        auth.signOut();
        goto("/login");
    }
</script>


<div class="banner">
  <nav>
	<a href="/">home</a>
	<a href="/">map</a>
    EMBARK
    <a href="#" on:click={logout}>Logout</a>
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

