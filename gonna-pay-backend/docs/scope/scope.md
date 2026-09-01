MVP Entities
USER: The account owner who authenticates and manages their data.
CONTACT: A private address book for the user. A contact created by User A is not visible to User B.
GROUP: A cost aggregator (e.g., "Beach Trip"). In the MVP, it is immutable (no editing, only deletion).
GROUP_MEMBERS: A join table linking CONTACTS to a GROUP. In the MVP, only the owner's contacts can be members.
COST: The primary entity (e.g., "Pizza $100"). It records who paid, the total amount, and the owner's share (ownerPercentage).
COST_SPLIT: Details how much each contact owes the owner. In the MVP, each contact gets their own split record (Option A).
Key Decisions and Changes (The "Why")
Removal of TRANSACTION: We eliminated the TRANSACTION entity to streamline the model. COST became the central entity because the app focuses on splitting specific expenses rather than managing complex banking workflows.
CONTACT as a Private Directory: We rejected the "ghost contact" idea (automatically linking contacts to real users). This keeps the MVP simple, avoids privacy issues, and focuses on immediate utility for the user.
One COST_SPLIT per Contact: We chose to create individual records for each debtor instead of a single group record. This ensures granularity for future reporting and individual debt tracking, with negligible impact on Postgres performance.
Equal Split in Groups: For costs linked to a group, the system automatically calculates the split between the owner and all members. This reduces the user's cognitive load during data entry.
No Editing (Immutability): To speed up MVP development, costs and groups cannot be edited. If a mistake occurs, the user deletes and recreates the entry. This drastically simplifies the Go backend logic.