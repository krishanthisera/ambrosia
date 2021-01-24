
CREATE TABLE stories(
    id              serial      NOT NULL,  
	name            varchar     PRIMARY KEY,
	description     varchar     NOT NULL,
	creation        varchar     NOT NULL,
	last_edit       varchar     NOT NULL,
	stories         varchar     NOT NULL
);