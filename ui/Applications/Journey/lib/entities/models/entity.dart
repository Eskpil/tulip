import 'package:equatable/equatable.dart';

class Entity extends Equatable {
  final String id;
  final String driver;

  final String deviceId;

  final Map<String, dynamic> entityMetadata;
  final Map<String, dynamic> driverMetadata;

  final String name;
  final String kind;

  const Entity({
    required this.id,
    required this.driver,

    required this.deviceId,

    required this.entityMetadata,
    required this.driverMetadata,

    required this.name,
    required this.kind,
  });


  factory Entity.fromJson(Map<String, dynamic> json) {
    return Entity(
      id: json['id'],
      driver: json['driver'],

      deviceId: json['device_id'],

      entityMetadata: json['entity_metadata'],
      driverMetadata: json['driver_metadata'],

      name: json['name'],
      kind: json['kind'],
    );
  }

  @override
  List<Object> get props => [id, driver, deviceId, entityMetadata, driverMetadata, name, kind];
}