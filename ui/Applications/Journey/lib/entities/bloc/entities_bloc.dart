import 'dart:async';
import 'package:bloc/bloc.dart';
import 'package:journey/entities/bloc/entities_event.dart';
import 'package:journey/entities/bloc/entities_state.dart';
import 'package:journey/repositories/entity.dart';

class EntitiesBloc extends Bloc<EntitiesEvent, EntitiesState> {
  final EntityRepository entityRepository;

  EntitiesBloc({ required this.entityRepository}) : super(EntitiesLoading()) {


    on<EntitiesStarted>(_onStarted);
  }

  Future<void> _onStarted(EntitiesStarted event, Emitter<EntitiesState> emit) async {
    emit(EntitiesLoading());
    try {
      final entities = await entityRepository.loadEntities();
      emit(EntitiesLoaded(entities: entities));
    } catch (_) {
      emit(EntitiesError());
    }
  }
}