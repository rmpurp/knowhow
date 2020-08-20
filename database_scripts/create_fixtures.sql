INSERT INTO pageContent VALUES ("git subtrees", "The Internet is full of articles on why you shouldn’t use Git submodules. While submodules are useful for a few use cases, they do have several drawbacks.");
INSERT INTO pageContent VALUES ("git subtrees", "git subtree lets you nest one repository inside another as a sub-directory. It is one of several ways Git projects can manage project dependencies.");
INSERT INTO pageContent VALUES ("git modules", "Git addresses this issue using submodules. Submodules allow you to keep a Git repository as a subdirectory of another Git repository. This lets you clone another repository into your project and keep your commits separate.");

INSERT INTO pages VALUES (NULL, NULL);
INSERT INTO pages VALUES (NULL, NULL);

INSERT INTO versions VALUES (null, 0, 1, 1);
INSERT INTO versions VALUES (null, 3600 * 24, 1, 2);
INSERT INTO versions VALUES (null, 3600 * 24 * 2, 2, 3);

UPDATE pages SET currentVersion = 2 where id = 1;
UPDATE pages SET currentVersion = 3 where id = 2;

