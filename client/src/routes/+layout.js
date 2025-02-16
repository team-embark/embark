import { redirect } from '@sveltejs/kit';
import { auth } from "../firebase_yippee";

export const load = ({url}) => {
    if (!auth.currentUser && url.pathname !== "/login") {
        return redirect(308, "/login");
    }
}
