import 'dart:async';
import 'dart:convert';
import 'package:format/format.dart';
import 'package:http/http.dart' as http;
import 'package:journey/core/models/state.dart';
import 'package:sprintf/sprintf.dart';

class Entity {
  final String id;
  final String driver;

  final String deviceId;

  final Map<String, dynamic> entityMetadata;
  final Map<String, dynamic> driverMetadata;

  final String name;
  final String kind;

  final List<EntityState> history;

  const Entity({
    required this.id,
    required this.driver,

    required this.deviceId,

    required this.entityMetadata,
    required this.driverMetadata,

    required this.name,
    required this.kind,

    required this.history,
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

      history: [],
    );
  }

  static Future<EntityState> state(Entity entity) async {
    final uri = Uri.parse(format("http://localhost:8000/entities/{}/history/last/", entity.id));
    final response = await http
        .get(uri);

    if (response.statusCode == 200) {
      // If the server did return a 200 OK response,
      // then parse the JSON.
      return EntityState.fromJson(jsonDecode(response.body));
    } else {
      // If the server did not return a 200 OK response,
      // then throw an exception.
      final code = response.statusCode;
      throw Exception('Failed to load entity state: $code');
    }
  }
}

Future<Entity> fetchEntity(String id) async {
  final uri = Uri.parse(sprintf("http://localhost:8000/entities/%s", id));
  final response = await http
    .get(uri);

  if (response.statusCode == 200) {
    // If the server did return a 200 OK response,
    // then parse the JSON.
    return Entity.fromJson(jsonDecode(response.body));
  } else {
    // If the server did not return a 200 OK response,
    // then throw an exception.
    throw Exception('Failed to load entity');
  }
}

Future<List<Entity>> fetchEntities() async {
  final response = await http
      .get(Uri.parse("http://localhost:8000/entities/"));

  if (response.statusCode == 200) {
    final entityList = jsonDecode(response.body);
    List<Entity> entities = [];

    for(final e in entityList){
      entities.add(Entity.fromJson(e));
    }

    return entities;
  } else {
    throw Exception("Failed to load entities");
  }
}