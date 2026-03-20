import { useState } from "react";
import { createLink, Link } from "../api";
import styles from "./ShortenModal.module.css";

interface ShortenModalProps {
	onClose: () => void;
	onCreated: (link: Link) => void;
}

export default function ShortenModal({ onClose, onCreated }: ShortenModalProps) {
	const [url, setUrl] = useState("");
	const [code, setCode] = useState("");
	const [error, setError] = useState<string | null>(null);
	const [loading, setLoading] = useState(false);

	async function submit() {
		if (!url) return;
		setLoading(true);
		setError(null);

		try {
			const response = await createLink(url, code || undefined);
			const data = await response.json();
			if (!response.ok) {
				setError(data.error ?? "Something went wrong.");
			} else {
				onCreated({ id: Date.now(), path: data.code, url });
			}
		} finally {
			setLoading(false);
		}
	}

	return (
		<div className={styles.overlay} onClick={onClose}>
			<div className={styles.modal} onClick={e => e.stopPropagation()}>
				<p className={styles.modalTitle}>Shorten a Link</p>
				{error && <div className={styles.modalError}>{error}</div>}
				<div className={styles.field}>
					<label className={styles.label}>URL</label>
					<input
						className={styles.input}
						placeholder="https://example.com/example"
						value={url}
						onChange={e => setUrl(e.target.value)}
						autoFocus
					/>
				</div>
				<div className={styles.field}>
					<label className={styles.label}>Custom code <span className={styles.labelOptional}>(optional)</span></label>
					<input
						className={styles.input}
						placeholder="my-link"
						value={code}
						onChange={e => setCode(e.target.value)}
					/>
				</div>
				<div className={styles.modalActions}>
					<button className={`${styles.button} ${styles.buttonCancel}`} onClick={onClose}>Cancel</button>
					<button className={`${styles.button} ${styles.buttonSubmit}`} onClick={submit} disabled={!url || loading}>
						{loading ? "Shortening..." : "Shorten"}
					</button>
				</div>
			</div>
		</div>
	);
}
