import {CredentialResponse} from "google-one-tap";

export interface opts {
    client_id: string,
    callback: (response: CredentialResponse) => void
}

export default function google_on_tap(opts: opts): void {
    const googleScript = document.createElement('script');
    googleScript.setAttribute('src', 'https://accounts.google.com/gsi/client');
    document.head.appendChild(googleScript)
    window.onload = function () {
        if (opts.client_id) {
            window.google.accounts.id.initialize({
                client_id: opts.client_id,
                callback: opts.callback,
                auto_select: false,
                cancel_on_tap_outside: true,
            });
            window.google.accounts.id.renderButton(document.getElementById('gsi-btn') as HTMLElement, {
                text: 'continue_with',
                locale: 'en'
            })
            window.google.accounts.id.prompt();
        } else {
            console.error('client_id is missing');
        }
    }
}