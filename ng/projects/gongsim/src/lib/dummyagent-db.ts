// insertion point for imports
import { EngineDB } from './engine-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class DummyAgentDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	TechName?: string
	Name?: string

	// insertion point for other declarations
	Engine?: EngineDB
	EngineID?: NullInt64

}
