import { useState } from "react";
import { inviteUser } from "../api";
import styles from "./InviteModal.module.css";

interface InviteModalProps {
	onClose: () => void;
}

export default function InviteModal({ onClose }: InviteModalProps) {
	const [email, setEmail] = useState("");
	const [error, setError] = useState<string | null>(null);
	const [success, setSuccess] = useState(false);
	const [loading, setLoading] = useState(false);

	async function submit() {
		if (!email) return;
		setLoading(true);
		setError(null);

		try {
			const response = await inviteUser(email);
			const data = await response.json();
			
			if (!response.ok) {
				setError(data.error ?? "Something went wrong.");
			} else {
				setSuccess(true);
			}
		} finally {
			setLoading(false);
		}
	}

	return (
		<div className={styles.overlay} onClick={onClose}>
			<div className={styles.modal} onClick={e => e.stopPropagation()}>
				<p className={styles.modalTitle}>Invite a User</p>
				{error && <div className={styles.modalError}>{error}</div>}
				{success ? (
					<>
						<p style={{ color: "#6aad6e", fontSize: "0.9rem" }}>
							{email} has been added to the whitelist.
						</p>
						<div className={styles.modalActions}>
							<button className={`${styles.button} ${styles.buttonSubmit}`} onClick={onClose}>Done</button>
						</div>
					</>
				) : (
					<>
						<div className={styles.field}>
							<label className={styles.label}>Email</label>
							<input
								className={styles.input}
								placeholder="example@example.com"
								type="email"
								value={email}
								onChange={e => setEmail(e.target.value)}
								autoFocus
							/>
						</div>
						<div className={styles.modalActions}>
							<button className={`${styles.button} ${styles.buttonCancel}`} onClick={onClose}>Cancel</button>
							<button className={`${styles.button} ${styles.buttonSubmit}`} onClick={submit} disabled={!email || loading}>
								{loading ? "Inviting..." : "Invite"}
							</button>
						</div>
					</>
				)}
			</div>
		</div>
	);
}
