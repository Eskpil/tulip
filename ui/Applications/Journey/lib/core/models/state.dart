class EntityState {
  final String entityId;
  final Map<String, dynamic> attributes;
  final Map<String, dynamic> state;

  const EntityState({
    required this.entityId,
    required this.attributes,
    required this.state,
  });

  factory EntityState.fromJson(Map<String, dynamic> json) {
    return EntityState(
      attributes: json["attributes"],
      entityId: json["entity_id"],
      state: json["state"],
    );
  }
}