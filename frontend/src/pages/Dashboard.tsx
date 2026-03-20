import { useEffect, useState } from "react";
import { getLinks, Link } from "../api";
import styles from "./Dashboard.module.css";
import InviteModal from "../components/InviteModal";
import ShortenModal from "../components/ShortenModal";
import LinkCard from "../components/LinkCard";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faRightFromBracket } from "@fortawesome/free-solid-svg-icons";

type ModalChoice = "shorten" | "invite" | null;

export default function Dashboard({ onLogout }: { onLogout: () => void }) {
	const [links, setLinks] = useState<Link[]>([]);
	const [modal, setModal] = useState<ModalChoice>(null);

	useEffect(() => {
		getLinks()
			.then(res => res.json())
			.then((data: { links: Link[] }) => setLinks(data.links ?? []));
	}, []);

	function handleLinkCreated(link: Link) {
		setLinks(prev => [...prev, link]);
		setModal(null);
	}

	function handleLinkDeleted(path: string) {
		setLinks(prev => prev.filter(l => l.path !== path));
	}

	return (
		<div className={styles.page}>
			<header className={styles.header}>
				<span className={styles.logo}>LinkFrog</span>
				<div className={styles.headerActions}>
					<button className={`${styles.button} ${styles.buttonInvite}`} onClick={() => setModal("invite")}>
						Invite User
					</button>
					<button className={`${styles.button} ${styles.buttonShorten}`} onClick={() => setModal("shorten")}>
						+ Shorten Link
					</button>
					<button className={`${styles.button} ${styles.buttonLogout}`} onClick={onLogout} title="Log out">
						<FontAwesomeIcon icon={faRightFromBracket} />
					</button>
				</div>
			</header>

			<main className={styles.content}>
				<p className={styles.sectionTitle}>Your Links</p>
				{links.length === 0 ? (
					<div className={styles.emptyState}>
						<p>No links yet. Create one to get started.</p>
					</div>
				) : (
					<div className={styles.linkList}>
						{links.map(link => (
							<LinkCard
								key={link.id}
								link={link}
								onDeleted={handleLinkDeleted}
							/>
						))}
					</div>
				)}
			</main>

			{modal === "shorten" && (
				<ShortenModal onClose={() => setModal(null)} onCreated={handleLinkCreated} />
			)}
			{modal === "invite" && (
				<InviteModal onClose={() => setModal(null)} />
			)}
		</div>
	);
}
