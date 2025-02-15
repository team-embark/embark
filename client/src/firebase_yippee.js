import { getAuth, GoogleAuthProvider} from "firebase/auth";
import { initializeApp } from 'firebase/app';

const firebaseConfig = {
    apiKey: "AIzaSyBC3lVKJ8BimNVOMGvkg4xQsG_mc8ow9QM",
    authDomain: "embark-6253e.firebaseapp.com",
    projectId: "embark-6253e",
    storageBucket: "embark-6253e.firebasestorage.app",
    messagingSenderId: "400488706588",
    appId: "1:400488706588:web:9fadf9f6eec4061b0a65d9"
};

export const app = initializeApp(firebaseConfig);
export const auth = getAuth();
export const provider = new GoogleAuthProvider();
