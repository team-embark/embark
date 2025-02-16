<script>
    import Button from "../components/Button.svelte";
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
	<Button class="transparentBtn" on:click={() => goto("/")}>
        <img src="/favicon.png" alt="logo" width= 30/>
    </Button>
    EMBARK
<<<<<<< HEAD
    <Button class="transparentBtn" on:click={logout}>
        Log Out
    </Button>
=======
    <a href="/signin">Sign Out</a>
    <a href="/Journal">journal</a>
>>>>>>> dde755b (Fixing something)
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

