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
            <img src="/embark-icon.png" alt="logo" width="50"/>
        </Button>
        <img src="/embark-text.png" alt="California2" width= "150" style="position: relative; left: 50px; top: 5px;z-index=11;"/>  
        <Button class="transparentBtn1" on:click={logout}>
            Log Out
        </Button>
    </nav>
</div>

<style>
    .banner {
        position:fixed;
        background:hsl(0, 0.00%, 100.00%);
        padding: 20px;
        width:95%;
        color: white;
        font-size: 1.5rem;
        border-radius: 4px;
        z-index:10;
        top: 0;
    }
</style>

{@render children()}

