# Slice order with maps

Map iteration order is random — cannot build ordered output by ranging a map. Correct pattern: one pass, map as seen-set, append to result on first sight. Also: append realloc copies; old backing array lives if other slices still reference it.
