<style>
    :global(body) {
      background-color: #b5e6ff;
    }
    .image-container {
        display: flex;
        flex-wrap: wrap;
        gap: 20px;
        position: relative; 
    }

    .right-panel {
        position: absolute;
        top: 120px; 
        right: 40px; 
        font-size: 5.0rem;
        font-weight: bold;
        text-align: right;
        background-color: rgba(255, 255, 255, 0.8); 
        padding: 10px;
        border-radius: 5px;
    }
  </style>

<script>
    import Button from "../components/Button.svelte"; 
    import {goto} from "$app/navigation";
    import {app, auth} from "../firebase_yippee";
    import {onAuthStateChanged, GoogleAuthProvider} from 'firebase/auth';
 

    let visitedParks = $state(0);
    const parks = ["josh", "yosemite", "redwood"];
    let data = $state(new Map());
    let disable = $state(new Map());

    $effect(()=> {
        onAuthStateChanged(auth, (user) => {
        if (user) {
            console.log("yay");
            countParks();
        } else {
            goto("/login");
        }
    });
        
    });

    async function countParks() {
        var myHeaders = new Headers();
        myHeaders.append("Authorization", "Bearer "+auth.currentUser.accessToken);
        var requestOptions = {
            method: 'GET',
            headers: myHeaders,
            redirect: 'follow'
        };
        const response = await fetch("http://localhost:8080/allparks", requestOptions)
            .then(response => response.json())
            .then(result => {
                for (const p of parks) {
                    if (result[p]){
                        visitedParks += 1;
                        data.set(p, 0);
                        disable.set(p, false);
                    } else {
                        data.set(p, 1);
                        disable.set(p, true);
                    }
                }
                console.log(disable);
                console.log(data);
            })
            .catch(error => console.log('error', error));
    }
</script>


<div class="image-container">
    <img src="/California2.png" alt="California2" width= "1300"/>  

<Button class="button1" on:click={() => {console.log("Yosemite");}} disabled={disable.get("yosemite")} style={`filter: grayscale(${data["yosemite"]});`}>

    <img src="/YosemiteNationalPark-2.png" alt="YosemiteNationalPark-2" width= "100"/>
    <br>
    Yosemite
</Button>

<Button class="button2" on:click={() => {console.log("Redwood");}} disabled={disable.get("redwood")} style={`filter: grayscale(${data["redwood"]});`}>
    <img src="/RedWoodNationalPark-2.png" alt="RedWoodNationalPark-2" width= "100" />
    <br>
    Redwood
</Button>

<Button class="button3" on:click={() => {goto("/journal/joshua-tree");}} disabled={disable.get("josh")} style={`filter: grayscale(${data["josh"]});`}>
    <img src="/JoshuaTree-2.png" alt="JoshuaTree-2" width= "100" />
    <br>
    Joshua Tree
</Button>
   
<Button class="button4" on:click={() => {{console.log("Channel Islands");}}} disabled={true}>
    <img src="/channelIslands.png" alt="channelIslands" width= "100" />
    <br>
    Channel Islands
</Button>

<Button class="button5" on:click={() => console.log("LassenVolcanic")} disabled={true}>
    <img src="/LassenVolcanic.png" alt="LassenVolcanic" width= "100" />
    <br>
    Lassen Volcanic
</Button>


<div class="right-panel">
    
    <p>{visitedParks}/5</p>
</div>

</div>
