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
        <img src="/embark-text.png" alt="California2" width= "150"/>  
        <Button class="transparentBtn" style="position: absolute; top: 1000px;" on:click={logout}>
            Log Out
        </Button>
    </nav>
</div>

<style>
    .banner {
        background:hsl(0, 0.00%, 100.00%);
        padding: 20px;
        color: white;
        font-size: 1.5rem;
        text-align: center;
        border-radius: 4px;
    }
</style>

{@render children()}

