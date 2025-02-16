<script>
    import Button from "./button.svelte"
    import {app, auth, provider} from '../../firebase_yippee';
    import {signInWithPopup, onAuthStateChanged, GoogleAuthProvider} from 'firebase/auth';
    import { redirect } from "@sveltejs/kit";

    let buttonText = "Sign in my scarab";

    function handleLogin() {
        if (!auth.currentUser) {
            signInWithPopup(auth, provider)
                .then(async (result) => {
                    // This gives you a Google Access Token. You can use it to access the Google API.
                    const credential = GoogleAuthProvider.credentialFromResult(result);
                    const token = credential.accessToken;
                    // The signed-in user info.
                    const user = result.user;
                    
                    const idToken = await user.getIdToken();
                    console.log(auth.currentUser,idToken);
                    sendToken(idToken);
                }).catch((error) => {
                    console.log("oh naur");
                    // Handle Errors here.
                    const errorCode = error.code;
                    const errorMessage = error.message;
                    console.log(errorCode,errorMessage);
                });
            // redirect(308, "/");           
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

</script>

<Button class = "primary sm" on:click={handleLogin}>
    {buttonText}
</Button>
