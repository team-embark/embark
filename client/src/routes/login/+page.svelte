<script>
    import {app, auth, provider} from '../../firebase_yippee';
    import {signInWithPopup, onAuthStateChanged, GoogleAuthProvider} from 'firebase/auth';
    import {goto} from "$app/navigation";

    let buttonText = "Log in";

    function handleLogin() {
        console.log("hi");
        if (!auth.currentUser) {
            signInWithPopup(auth, provider)
                .then(async (result) => {
                    // This gives you a Google Access Token. You can use it to access the Google API.
                    const credential = GoogleAuthProvider.credentialFromResult(result);
                    const token = credential.accessToken;
                    console.log(token);
                    // The signed-in user info.
                    const user = result.user;
                    
                    const idToken = await user.getIdToken();
                    console.log(auth.currentUser,idToken);
                    sendToken(idToken);
                    goto("/");
                }).catch((error) => {
                    console.log("oh naur");
                    // Handle Errors here.
                    const errorCode = error.code;
                    const errorMessage = error.message;
                    console.log(errorCode,errorMessage);
                });
            // redirect(308, "/");           
        } else {
            auth.signOut();
        }
    }

    async function sendToken(token) {
        var myHeaders = new Headers();
        myHeaders.append("Authorization", "Bearer "+token);

        var requestOptions = {
            method: 'POST',
            headers: myHeaders,
            redirect: 'follow'
        };
        const response = await fetch("http://localhost:8080/auth", requestOptions)
            .then(response => response.text())
            .then(result => console.log(result))
            .catch(error => console.log('error', error));
    }
    onAuthStateChanged(auth, (user) => {
        if(user) {
            buttonText = "Log out";
        }
        else {
            buttonText = "Log in";
        }
    })
</script>


<img src="/backgroundLogin.png" alt="/backgroundLogin" width= "1500" style="position: absolute; left: 0px; top: 0px;z-index=11;"/>
<img src="/iconwithwhite.png" alt="/iconwithwhite" width= "100" style="position: fixed; left: 650px; top: 200px;z-index=17;"/>
<div class="container">
    <button class="signin" style="position: fixed; left: 600px; top: 300px; z-index: 100;" on:click={handleLogin}>
        {buttonText}
    </button>
</div>

<style>
    div.container {
        display:flex;
        justify-content: center;
        align-items: center;
        height: 60vh;
    }
    .signin {
        border: none;
        border-radius: 20px;
        width: 200px;
        height: 60px;
        font-size: 25px;
        font-family: "Futura";
    }
    .signin:hover {
        transform: scale(1, 1.01);
        background-color: rgb(214, 228, 221);
    }

</style>
