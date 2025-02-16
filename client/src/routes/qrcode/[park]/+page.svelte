<script>
    import { onAuthStateChanged } from "firebase/auth";
    import {auth} from "../../../firebase_yippee";
    import { page } from "$app/stores";
    import CenteredBox from "./CenteredBox.svelte";
    import NewTextComponent from "./NewTestComponent.svelte";
    import ButtonComponent from "./ButtonComponent.svelte";
    import { goto } from "$app/navigation";


    let parkname = $state("");
    let parks = {
        "yosemite": "Yosemite",
        "josh": "Joshua Tree",
        "redwood" : "Redwood"
    }

    function isBase64(str) {
        try {
            return btoa(atob(str)) === str;
        } catch (err) {
            return false;
        }
    }
    
    $effect(() => {
        let park64 = $page.params.park;
        onAuthStateChanged(auth, (user) => {
            console.log(park64);
            if (user && isBase64(park64)) {
                parkname = atob(park64);
                console.log(parkname);
                visitPark();
            } else {
                parkname = "oops";
            }
        });
    });

    async function visitPark() {
        var myHeaders = new Headers();
        myHeaders.append("Authorization", "Bearer "+auth.currentUser.accessToken);

        var requestOptions = {
            method: 'POST',
            headers: myHeaders,
            redirect: 'follow'
        };
        const response = await fetch(`http://localhost:8080/parks/${parkname}`, requestOptions)
            .then(response => response.text())
            .then(result => console.log(result))
            .catch(error => console.log('error', error));
    }

</script>

<div>
	<p>Thanks for visiting {parks[parkname]}!</p>
	<!-- <CenteredBox width="60em" height="40em" /> -->
	<!-- <NewTextComponent text="Fun fact!" /> -->
     <button on:click={()=>{goto("/journal/joshua-tree")}}>Go to Journal</button>
	<img src="/YosemiteNationalPark-2.png" alt="YosemiteNationalPark-2" width= "200" class="image"/>
</div>

<style>
	p {
		color: goldenrod;
		font-family: 'Futura';
		font-size: 2em;
        text-align: center;
        margin-top: 150px;
	}
    .image {
		position: fixed;
		bottom: 20%;
		right: 20%;
	}
</style>
