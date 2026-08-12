# Value vs pointer receiver

Learner can use pointer receivers for mutation and asked why Balance should be value. Rule internalized: mutate → pointer; read-only → value (unless size/mutex/consistency). Method sets: value methods callable on *T via auto-deref.
