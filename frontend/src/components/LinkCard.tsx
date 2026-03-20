import { useState } from "react";
import { deleteLink, Link } from "../api";
import styles from "./LinkCard.module.css";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faCopy, faCheck, faTrash } from "@fortawesome/free-solid-svg-icons";

interface LinkCardProps {
	link: Link;
	onDeleted: (path: string) => void;
}

export default function LinkCard({ link, onDeleted }: LinkCardProps) {
	const [copied, setCopied] = useState(false);

	function copy() {
		navigator.clipboard.writeText(`${window.location.origin}/l/${link.path}`);
		setCopied(true);
		setTimeout(() => setCopied(false), 1500);
	}

	async function handleDelete() {
		const response = await deleteLink(link.path);
		if (response.ok) {
			onDeleted(link.path);
		}
	}

	return (
		<div className={styles.linkCard}>
			<div className={styles.linkInfo}>
				<div className={styles.linkCode}>/{link.path}</div>
				<div className={styles.linkUrl}>{link.url}</div>
			</div>
			<div className={styles.linkActions}>
				<button className={styles.buttonIcon} onClick={copy}>
					<FontAwesomeIcon icon={copied ? faCheck : faCopy} />
				</button>
				<button className={`${styles.buttonIcon} ${styles.danger}`} onClick={handleDelete}>
					<FontAwesomeIcon icon={faTrash} />
				</button>
			</div>
		</div>
	);
}
