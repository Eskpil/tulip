import 'dart:convert';
import 'package:format/format.dart';
import 'package:http/http.dart' as http;

import 'package:web_socket_channel/web_socket_channel.dart';
import 'package:journey/core/models/entity.dart';
import 'package:journey/core/models/state.dart';

class EntityRepository {
  final WebSocketChannel channel =
  WebSocketChannel.connect(Uri.parse("ws://localhost:8004/gateway/"));


  Future<List<Entity>> loadEntities() async {
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

  Future<EntityState> loadLatestEntityState(Entity entity) async {
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