import GoogleButton from "react-google-button";
import styles from "./LoginPage.module.css";

interface Props {
    unauthorized?: boolean
}

export default function LoginPage({ unauthorized }: Props) {
    return (
        <div className={styles.page}>
            <div className={styles.card}>
                <h1 className={styles.title}>LinkFrog</h1>
                <p className={styles.tagline}>Ribbit! Got a link?</p>
                {unauthorized && (
                    <div className={styles.error}>
                        You are not authorized to access this application.
                    </div>
                )}
                
                <GoogleButton style={{
                    marginLeft: "auto",
                    marginRight: "auto"
                }} onClick={() => { window.location.href = "/api/auth/google"; }} />
            </div>
        </div>
    );
}
