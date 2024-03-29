import 'dart:convert';

import 'package:bloc/bloc.dart';

import 'package:journey/entity/bloc/entity_state_event.dart';
import 'package:journey/entity/bloc/entity_state_state.dart';

import 'package:journey/core/models/entity.dart';
import 'package:journey/core/models/state.dart';

import '../../repositories/entity.dart';

class Event {
  final String subject;
  final EntityState state;

  const Event({ required this.subject, required this.state });

  factory Event.fromJson(Map<String, dynamic> json) {
    return Event(
      subject: json["subject"],
      state: EntityState.fromJson(json["payload"]),
    );
  }
}

class EntityStateBloc extends Bloc<EntityStateEvent, EntityStateState> {
  final EntityRepository entityRepository;
  final Entity entity;

  EntityStateBloc({ required this.entityRepository, required this.entity }) : super(EntityStateLoading()) {
    // FIXME: This should probably be behind some kind of event.
    entityRepository.channel.stream.forEach((rawState) {
      final json = jsonDecode(rawState);

      final event = Event.fromJson(json);

      if (event.subject == "state" && event.state.entityId == entity.id) {
        emit(EntityStateUpdated(state: event.state));
      }
    });

    on<EntityStateStarted>(_onStarted);
  }

  Future<void> _onStarted(EntityStateStarted event, Emitter<EntityStateState> emit) async {
    emit(EntityStateLoading());
    try {
      final state = await entityRepository.loadLatestEntityState(entity);
      emit(EntityStateUpdated(state: state));
    } catch (_) {
      emit(EntityStateError());
    }
  }
}
